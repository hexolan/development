import type { Post } from '../app/types/common'

const FeedPost = ({ post }: { post: Post }) => (
  <div>
    <h1>Title: {post.title}</h1>
    <span>Panel: {post.panelId} | Created by: {post.authorId}</span>
    <p>{post.content}</p>
  </div>
)

export default FeedPost