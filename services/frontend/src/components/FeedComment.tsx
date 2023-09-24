import { Link } from 'react-router-dom'
import { Paper, Group, Box, ThemeIcon, Text } from '@mantine/core'
import { IconMessage } from '@tabler/icons-react'

import { useGetUserByIdQuery } from '../app/api/users'
import type { Comment } from "../app/types/common"

const FeedComment = ({ comment }: { comment: Comment }) => {
  // fetching comment author info
  const { data } = useGetUserByIdQuery({ id: comment.authorId })
  let authorElement = <Text color='dimmed' size='xs'>Loading Author Info...</Text>
  if (!data) {
    authorElement = <Text color='red' size='xs'>Failed to load Author Info</Text>
  } else {
    authorElement = <Text color='dimmed' size='xs' mt={3} component={Link} to={`/user/${data.username}`}>by user/{data.username}</Text>
  }

  return (
    <Paper shadow='sm' radius='md' p='md' withBorder>
      <Group>
        <ThemeIcon color='teal' variant='light' size='xl'><IconMessage /></ThemeIcon>
        <Box>
          <Text size='sm'>{comment.message}</Text>
          {authorElement}
        </Box>
      </Group>
    </Paper>
  )
}

export default FeedComment