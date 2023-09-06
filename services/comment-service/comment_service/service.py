from typing import Type, List

from comment_service.models import CommentRepository, Comment, CommentCreate, CommentUpdate


class ServiceRepository(CommentRepository):
    def __init__(self, downstream_repo: Type[CommentRepository]) -> None:
        self._repo = downstream_repo

    async def get_post_comments(self, post_id: str) -> List[Comment]:
        # todo: pagination
        return await self._repo.get_post_comments(post_id)
    
    async def create_comment(self, data: CommentCreate) -> Comment:
        return await self._repo.create_comment(data)
    
    async def update_comment(self, comment_id: int, data: CommentUpdate) -> Comment:
        return await self._repo.update_comment(comment_id, data)
    
    async def delete_comment(self, comment_id: int) -> None:
        await self._repo.delete_comment(comment_id)