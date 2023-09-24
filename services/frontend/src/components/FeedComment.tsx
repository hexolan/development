import { Paper, Group, Box, ThemeIcon, Text } from '@mantine/core'
import { IconMessage } from '@tabler/icons-react'

import type { Comment } from "../app/types/common"

const FeedComment = ({ comment }: { comment: Comment }) => (
  // todo: fetching comment author
  <Paper shadow='sm' radius='md' p='md' withBorder>
    <Group>
      <ThemeIcon color='teal' variant='light' size='xl'><IconMessage /></ThemeIcon>
      <Box>
        <Text size='sm'>{comment.message}</Text>
        <Text color='dimmed' size='xs' mt={3}>by user/{comment.authorId}</Text>
      </Box>
    </Group>
  </Paper>
)

export default FeedComment