import hmac
import hashlib
from typing import Type

import jwt
from argon2 import PasswordHasher
from argon2.exceptions import VerificationError

from auth_service.models.config import Config
from auth_service.models.exceptions import ServiceException, ServiceErrorCode
from auth_service.models.service import AuthRepository, AuthDBRepository, AuthToken, AccessTokenClaims


class ServiceRepository(AuthRepository):
    def __init__(self, config: Config, downstream_repo: Type[AuthDBRepository]) -> None:
        self._repo = downstream_repo

        self._jwt_private = config.jwt_private_key

        self._hasher = PasswordHasher()  # Use default RFC_9106_LOW_MEMORY profile
        self._pepper = bytes(config.password_pepper, encoding="utf-8")

    def _apply_pepper(self, password: str) -> str:
        return hmac.new(key=self._pepper, msg=bytes(password, encoding="utf-8"), digestmod=hashlib.sha256).hexdigest()

    def _hash_password(self, password: str) -> str:
        return self._hasher.hash(self._apply_pepper(password))

    def _verify_password(self, hash: str, password: str) -> None:
        self._hasher.verify(hash, self._apply_pepper(password))

    def _issue_auth_token(self, user_id: str) -> AuthToken:
        claims = AccessTokenClaims(sub=user_id)
        access_token = jwt.encode(claims.model_dump(), self._jwt_private, algorithm="RS256")
        return AuthToken(access_token=access_token)

    async def auth_with_password(self, user_id: str, password: str) -> AuthToken:
        # Get the auth record for that user
        auth_record = await self._repo.get_auth_record(user_id)
        if not auth_record:
            raise ServiceException("invalid user id or password", ServiceErrorCode.INVALID_CREDENTIALS)

        # Verify the password hashes
        try:
            self._verify_password(auth_record.password.get_secret_value(), password)
        except VerificationError:
            raise ServiceException("invalid user id or password", ServiceErrorCode.INVALID_CREDENTIALS)
        
        # Update the auth record if the password needs rehash
        if self._hasher.check_needs_rehash(auth_record.password.get_secret_value()):
            await self._repo.update_password_auth_method(auth_record.user_id, self._hash_password(password))

        # Issue a token for the user
        return self._issue_auth_token(auth_record.user_id)

    async def set_password_auth_method(self, user_id: str, password: str) -> None:
        # Hash the password
        password = self._hash_password(password)

        # Update the auth method or create one if it doesn't exist
        auth_record = await self._repo.get_auth_record(user_id)
        if auth_record is not None:
            await self._repo.update_password_auth_method(user_id, password)
        else:
            await self._repo.create_password_auth_method(user_id, password)

    async def delete_password_auth_method(self, user_id: str) -> None:
        await self._repo.delete_password_auth_method(user_id)