import FeedPost from './FeedPost';
import type { Post } from '../app/types/common'

function Feed() {
  const testPost = {
    id: 'testing',
    panelId: 'test',
    authorId: 'test',
    title: 'test',
    content: 'test',
    createdAt: 'test'
  } as Post
  // todo: remove after

  return (
    <div>
      <FeedPost post={testPost} />
    </div>
  )
}

export default Feed