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
  const { data, isLoading } = useGetPanelPostQuery({ panelId: panel.id, postId: postId })
  if (isLoading) {
    return <LoadingBar />;
  } else if (!data) {
    throw Error('Post not found!')  // todo: extract exact error msg (as it may not just be a 404)
  }

  return <PagePost post={data} />
}

export default PanelPostPage