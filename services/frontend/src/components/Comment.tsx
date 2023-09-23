import type { Comment } from "../app/types/common"

const Comment = ({ comment }: { comment: Comment }) => (
  <div>
    <p>{comment.message}</p>
  </div>
)

export default Comment