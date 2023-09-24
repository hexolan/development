import { Link } from 'react-router-dom'
import { Paper, Group, Box, ThemeIcon, Text, ActionIcon, Menu } from '@mantine/core'
import { IconMessage, IconMenu2, IconTrash, IconPencil } from '@tabler/icons-react'

import { useAppSelector } from '../app/hooks'
import { useGetUserByIdQuery } from '../app/api/users'
import type { Comment } from "../app/types/common"

const FeedComment = ({ comment }: { comment: Comment }) => {
  const currentUser = useAppSelector((state) => state.auth.currentUser)

  // fetching comment author info
  const { data } = useGetUserByIdQuery({ id: comment.authorId })
  let authorElement = <Text color='dimmed' size='xs'>Loading Author Info...</Text>
  if (!data) {
    authorElement = <Text color='red' size='xs'>Failed to load Author Info</Text>
  } else {
    authorElement = <Text color='dimmed' size='xs' mt={3} component={Link} to={`/user/${data.username}`}>by user/{data.username}</Text>
  }

  // todo: add functionality for 'Modify' and 'Delete' comment buttons
  return (
    <Paper shadow='sm' radius='md' p='md' withBorder>
      <Group position='apart'>
        <Group>
          <ThemeIcon color='teal' variant='light' size='xl'><IconMessage /></ThemeIcon>
          <Box>
            <Text size='sm'>{comment.message}</Text>
            {authorElement}
          </Box>
        </Group>
        {currentUser && (currentUser.id == comment.authorId || currentUser.isAdmin) && (
          <Menu>
            <Menu.Target>
              <ActionIcon color='teal' variant='light' radius='xl' size='xl'><IconMenu2 /></ActionIcon>
            </Menu.Target>
            <Menu.Dropdown>
              <Menu.Label>Comment Options</Menu.Label>
              { currentUser.id == comment.authorId && <Menu.Item icon={<IconPencil size={14} />}>Modify</Menu.Item> }
              <Menu.Item color='red' icon={<IconTrash size={14} />}>Delete</Menu.Item>
            </Menu.Dropdown>
          </Menu>
        )}
      </Group>
    </Paper>
  )
}

export default FeedComment