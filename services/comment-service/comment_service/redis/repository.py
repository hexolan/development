import pickle
import logging
from typing import Type, List

import redis.asyncio as redis

from comment_service.models import CommentRepository, Comment, CommentCreate, CommentUpdate


class ServiceRedisRepository(CommentRepository):
    def __init__(self, redis_conn: redis.Redis, downstream_repo: Type[CommentRepository]) -> None:
        self._conn = redis_conn
        self._repo = downstream_repo

    async def _purge_cached_comments(self, post_id: str) -> None:
        try:
            self._conn.delete(post_id)
        except Exception:
            pass

    async def get_post_comments(self, post_id: str) -> List[Comment]:
        post_id = post_id.lower()

        failure = False
        try:
            # check for a cached version of the post comments
            response = await self._conn.get(post_id)
            if response is not None:
                comments = pickle.loads(response)
                assert type(comments) == list
                return comments
        except Exception as e:
            failure = True
            logging.error("Redis Repo: error whilst getting post comments", e)
            pass

        comments = await self._repo.get_post_comments(post_id)
        if failure is False:
            # cache the retrieved comments            
            try:
                await self._conn.set(post_id, pickle.dumps(comments), ex=120)  # TTL: 2 minutes
            except Exception:
                logging.error("Redis Repo: error whilst caching post comments", e)
                pass

        return comments
    
    async def create_comment(self, data: CommentCreate) -> Comment:
        comment = await self._repo.create_comment(data)
        await self._purge_cached_comments(comment.post_id)
        return comment
    
    async def update_comment(self, comment_id: int, data: CommentUpdate) -> Comment:
        comment = await self._repo.update_comment(comment_id, data)
        await self._purge_cached_comments(comment.post_id)
        return comment
    
    async def delete_comment(self, comment_id: int) -> None:
        # todo: purge cache of comments on post (instead of waiting for TTL expiry)
        await self._repo.delete_comment(comment_id)