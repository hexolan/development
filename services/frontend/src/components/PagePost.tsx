import { Link, useNavigate } from 'react-router-dom'
import { Paper, Stack, Badge, ThemeIcon, Text, Group, Menu, ActionIcon } from '@mantine/core'
import { IconUser, IconMenu2, IconTrash } from '@tabler/icons-react'

import { useGetUserByIdQuery } from "../app/api/users"
import { Post } from "../app/types/common"
import { useDeletePostMutation } from '../app/api/posts'

const PagePost = ({ post }: { post: Post }) => {
  const navigate = useNavigate()
  const { data } = useGetUserByIdQuery({ id: post.authorId })

  const [deletePost] = useDeletePostMutation()
  const submitDeletePost = async () => {
    // todo: change type to just take id: (to auto invalidate post by id - or experiment with RTK query tags)
    await deletePost({postId: post.id}).unwrap().then(() => {
      navigate('/')
    }).catch((error) => {
      console.log(error)  // todo: error handling
    })
  }

  return (
    <Paper shadow='lg' radius='lg' p='lg' withBorder>
      {data && (
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
            to={`/user/${data.username}`}
          >
            {`user/${data.username}`}
          </Badge>
          {/* todo: hiding menu when not admin / not post author (+ add functionality for Deleting and Updating)*/}
          <Menu>
            <Menu.Target>
              <ActionIcon color='teal' variant='light' radius='xl' size={24}><IconMenu2 size={12} /></ActionIcon>
            </Menu.Target>
            <Menu.Dropdown>
              <Menu.Label>Post Options</Menu.Label>
              <Menu.Item color='red' icon={<IconTrash size={14} />} onClick={() => submitDeletePost()}>Delete</Menu.Item>
            </Menu.Dropdown>
          </Menu>
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