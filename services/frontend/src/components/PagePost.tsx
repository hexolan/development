import { Link, useNavigate } from 'react-router-dom'
import { Paper, Stack, Badge, ThemeIcon, Text, Group, Menu, ActionIcon } from '@mantine/core'
import { IconUser, IconMenu2, IconPencil, IconTrash } from '@tabler/icons-react'

import { useAppSelector } from '../app/hooks'
import { useGetUserByIdQuery } from '../app/api/users'
import { useDeletePostMutation } from '../app/api/posts'
import type { Post } from '../app/types/common'

const PagePost = ({ post }: { post: Post }) => {
  const navigate = useNavigate()

  const currentUser = useAppSelector((state) => state.auth.currentUser)
  
  const [deletePost] = useDeletePostMutation()
  const submitDeletePost = async () => {
    // todo: change type to just take id: (to auto invalidate post by id - or experiment with RTK query tags)
    await deletePost({ id: post.id }).unwrap().then(() => {
      navigate('/')
    }).catch((error) => {
      console.log(error)  // todo: error handling
    })
  }
  
  const { data: authorData } = useGetUserByIdQuery({ id: post.authorId })
  return (
    <Paper shadow='lg' radius='lg' p='lg' withBorder>
      {authorData && (
        <Group position='apart'>
          <Badge
            pl={0}
            color='teal'
            leftSection={
              <ThemeIcon color='teal' size={24} radius='xl' mr={5}>
                <IconUser size={12} />
              </ThemeIcon>
            }
            component={Link}
            to={`/user/${authorData.username}`}
          >
            {`user/${authorData.username}`}
          </Badge>
          {(currentUser && (currentUser.id == post.authorId || currentUser.isAdmin)) && (
            <Menu>
              <Menu.Target>
                <ActionIcon color='teal' variant='light' radius='xl' size={24}>
                  <IconMenu2 size={12} />
                </ActionIcon>
              </Menu.Target>
              <Menu.Dropdown>
                <Menu.Label>Post Options</Menu.Label>
                { currentUser.id == post.authorId && <Menu.Item icon={<IconPencil size={14} />}>Modify</Menu.Item> }
                <Menu.Item color='red' icon={<IconTrash size={14} />} onClick={() => submitDeletePost()}>Delete</Menu.Item>
              </Menu.Dropdown>
            </Menu>
          )}
        </Group>
      )}
      <Stack align='flex-start' mt={2} spacing={1}>
        <Text weight={600}>{post.title}</Text>
        <Text size='sm'>{post.content}</Text>
        <Text size='xs' color='dimmed' mt={3}>Created {post.createdAt}</Text>
      </Stack>
    </Paper>
  )
}

export default PagePost