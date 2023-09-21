import FeedPost from './FeedPost'
import { useGetPanelPostsQuery } from '../app/api/posts'

function PanelPostFeed({ panelId }: { panelId: string }) {
  const { data, isLoading } = useGetPanelPostsQuery({ panelId: panelId })
  
  // todo: improve
  if (isLoading) {
    return <p>loading....</p>
  } else if (!data) {
    return <p>failed...</p>
  }

  if (!data.length) {
    return <p>no posts</p>
  } else {
    return (
      <>
        {Object.values(data).map(post => {
          return <FeedPost key={post.id} post={post} hidePanel={true} />
        })}
      </>
    )
  }
}

export default PanelPostFeed