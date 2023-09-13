import FeedPost from './FeedPost';
import type { Post } from '../app/types'

function Feed() {
  let testPost = {
    id: 'testing',
    panelId: 'test',
    authorId: 'test',
    title: 'test',
    content: 'test',
    createdAt: 'test'
  } as Post

  return (
    <div>
      <FeedPost post={testPost} />
    </div>
  )
}

export default Feed