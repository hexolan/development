import { useState } from 'react'
import { Paper, Flex, Textarea, ActionIcon } from '@mantine/core'
import { useForm, hasLength } from '@mantine/form'
import { IconWriting } from '@tabler/icons-react'

import { useCreatePostCommentMutation } from '../app/api/comments'
import type { Comment, Post } from '../app/types/common'
import type { CreateCommentData } from '../app/types/comments'

const CreateComment = ({ post, addNewComment }: { post: Post, addNewComment: (comment: Comment) => void }) => {
  const [errorMsg, setErrorMsg] = useState('')

  const [createComment, { isLoading }] = useCreatePostCommentMutation()
  const submitCommentForm = async (values: CreateCommentData) => {
    const commentInfo = await createComment({
      postId: post.id,
      data: values
    }).unwrap().catch(
      (error) => {
        if (!error.data) {
          setErrorMsg('Failed to access the API')
        } else {
          setErrorMsg(error.data.msg)
        }
      }
    )

    // display the new comment
    if (commentInfo) {
      addNewComment(commentInfo)
    }
  }

  const commentForm = useForm<CreateCommentData>({
    initialValues: {
      message: '',
    },
    validate: {
      message: hasLength({ min: 3, max: 2048 }, 'Message must be between 3 and 2048 characters'),
    }
  })

  return (
    <Paper shadow='sm' radius='md' p='md' withBorder>
      <form onSubmit={commentForm.onSubmit(submitCommentForm)}>
        <Flex gap='sm' align='center' direction='row' wrap='nowrap'>
            <Textarea 
              size='xs'
              w='100%'
              radius='lg'
              variant='filled'
              placeholder='Input comment...'
              error={errorMsg}
              {...commentForm.getInputProps('message')}
            />
            
            <ActionIcon type='submit' radius='lg' color='teal' variant='outline' size='xl' aria-label='Post Comment' disabled={isLoading}>
              <IconWriting />
            </ActionIcon>
        </Flex>
      </form>
    </Paper>
  )
}

export default CreateComment