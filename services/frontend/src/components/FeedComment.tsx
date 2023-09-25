import { useState } from 'react'
import { Link } from 'react-router-dom'
import { useForm, hasLength } from '@mantine/form'
import { Paper, Group, Box, ThemeIcon, Text, ActionIcon, Menu, Textarea, Flex } from '@mantine/core'
import { IconMessage, IconMenu2, IconTrash, IconPencil, IconPencilCancel } from '@tabler/icons-react'

import { useAppSelector } from '../app/hooks'
import { useGetUserByIdQuery } from '../app/api/users'
import { useUpdatePostCommentMutation } from '../app/api/comments'
import type { Comment } from "../app/types/common"
import type { UpdateCommentData } from '../app/types/comments'

const FeedComment = ({ comment: providedComment }: { comment: Comment }) => {
  const [comment, setComment] = useState<Comment>(providedComment)
  const currentUser = useAppSelector((state) => state.auth.currentUser)

  const [modifying, setModifying] = useState<boolean>(false)
  const [errorMsg, setErrorMsg] = useState('')
  const commentForm = useForm<UpdateCommentData>({
    initialValues: {
      message: comment.message,
    },
    validate: {
      message: hasLength({ min: 3, max: 512 }, 'Message must be between 3 and 512 characters'),
    }
  })
  
  const [updateComment, { isLoading }] = useUpdatePostCommentMutation()
  const submitCommentForm = async (values: UpdateCommentData) => {
    const commentInfo = await updateComment({
      id: comment.id,
      postId: comment.postId,
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

    // display the updated comment
    if (commentInfo) {
      setComment(commentInfo)
      setModifying(false)
    }
  }

  // fetching comment author info
  const { data } = useGetUserByIdQuery({ id: comment.authorId })
  let authorElement = <Text color='dimmed' size='xs'>Loading Author Info...</Text>
  if (!data) {
    authorElement = <Text color='red' size='xs'>Failed to load Author Info</Text>
  } else {
    authorElement = <Text color='dimmed' size='xs' mt={3} component={Link} to={`/user/${data.username}`}>by user/{data.username}</Text>
  }

  // todo: add functionality for 'Delete' comment button
  // improve layout of editing comment (fix flexboxes)
  return (
    <Paper shadow='sm' radius='md' p='md' withBorder>
      <Flex gap='sm' align='center' direction='row' wrap='nowrap'>
        <Group w='100%'>
          <ThemeIcon color='teal' variant='light' size='xl'><IconMessage /></ThemeIcon>
          {modifying ? (
            <Box w='90%'>
              <form onSubmit={commentForm.onSubmit(submitCommentForm)}>
                <Textarea size='xs' w='100%' radius='lg' variant='filled' error={errorMsg} {...commentForm.getInputProps('message')} />
                <ActionIcon type='submit' radius='lg' color='teal' variant='outline' size='xl' aria-label='Update Comment' disabled={isLoading}>
                  <IconPencil />
                </ActionIcon>
              </form>
            </Box>
          ) : (
            <Box>
              <Text size='sm'>{comment.message}</Text>
              {authorElement}
            </Box>
          )}
        </Group>
        {currentUser && (currentUser.id == comment.authorId || currentUser.isAdmin) && (
          <Menu>
            <Menu.Target>
              <ActionIcon color='teal' variant='light' radius='xl' size='xl'><IconMenu2 /></ActionIcon>
            </Menu.Target>
            <Menu.Dropdown>
              <Menu.Label>Comment Options</Menu.Label>
              {currentUser.id == comment.authorId && (
                modifying ? <Menu.Item icon={<IconPencilCancel size={14} />} onClick={() => setModifying(false)}>Stop Modifying</Menu.Item>
                : <Menu.Item icon={<IconPencil size={14} />} onClick={() => setModifying(true)}>Modify</Menu.Item>
              )}
              <Menu.Item color='red' icon={<IconTrash size={14} />}>Delete</Menu.Item>
            </Menu.Dropdown>
          </Menu>
        )}
      </Flex>
    </Paper>
  )
}

export default FeedComment