import { useParams, useOutletContext } from 'react-router-dom'

import PagePost from '../components/PagePost'
import LoadingBar from '../components/LoadingBar'
import { useGetPanelPostQuery } from '../app/api/posts'
import type { PanelContext } from '../components/PanelLayout'

type PanelPostPageParams = {
  panelName: string;
  postId: string;
}

function PanelPostPage() {
  const { panel } = useOutletContext<PanelContext>()
  const { postId } = useParams<PanelPostPageParams>();
  if (postId === undefined) { throw Error('post id not provided') }

  // Fetch the post
  const { data, error, isLoading } = useGetPanelPostQuery({ panelId: panel.id, postId: postId })
  if (isLoading) {
    return <LoadingBar />;
  } else if (!data) {
    if (!error) {
      throw Error('Unknown error occured')
    } else if (!error.data) {
      throw Error('Failed to access the API')
    } else {
      throw Error(error.data.msg)
    }
  }

  // todo: PostComments
  return <PagePost post={data} />
}

export default PanelPostPage