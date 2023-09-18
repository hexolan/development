import { Paper, Stack, Avatar, Badge, Text } from '@mantine/core';

import type { Post } from '../app/types/common'

const FeedPost = ({ post }: { post: Post }) => {
  const panelName = "todoFetchPanel"  // post.panelId
  const authorName = "todoFetchAuthor"  // post.authorId

  return (
    <Paper shadow="xl" radius="lg" p="lg" withBorder>
      <Stack align="flex-start" spacing={1}>
        <Badge color='teal' pl={0} leftSection={<Avatar size={24} color='green' radius='xl' mr={5} alt={'Created by ' + authorName} />}>{'user/' + authorName}</Badge>
        <Text>{post.title}</Text>
        <Text color='dimmed' size='xs'>{'panel/' + panelName}</Text>
        <Text size='sm' truncate>{post.content}</Text>
      </Stack>
    </Paper>
  )
}

export default FeedPost