from typing import Any, List

from pydantic import computed_field
from pydantic.fields import FieldInfo
from pydantic_settings import BaseSettings, EnvSettingsSource
from pydantic_settings.main import BaseSettings
from pydantic_settings.sources import PydanticBaseSettingsSource


class ConfigSource(EnvSettingsSource):
    def prepare_field_value(self, field_name: str, field: FieldInfo, value: Any, value_is_complex: bool) -> Any:
        if field_name == "kafka_brokers":
            if value == None:
                return None
            return value.split(",")

        return super().prepare_field_value(field_name, field, value, value_is_complex)


class Config(BaseSettings):
    postgres_user: str
    postgres_pass: str
    postgres_host: str
    postgres_database: str

    redis_host: str
    redis_pass: str

    kafka_brokers: List[str]

    @computed_field
    @property
    def postgres_dsn(self) -> str:
        return "postgresql+asyncpg://{user}:{password}@{host}/{db}".format(
            user=self.postgres_user,
            password=self.postgres_pass,
            host=self.postgres_host,
            db=self.postgres_database
        )
    
    @classmethod
    def settings_customise_sources(cls, settings_cls: type[BaseSettings], *args, **kwargs) -> tuple[PydanticBaseSettingsSource, ...]:
        return (ConfigSource(settings_cls), )