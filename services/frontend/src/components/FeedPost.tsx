import { Paper, Stack, Avatar, Badge, Text } from '@mantine/core';

import type { Post } from '../app/types/common'

const FeedPost = ({ post }: { post: Post }) => (
  <Paper shadow="xl" radius="lg" p="lg" withBorder>
    <Stack align="flex-start" spacing={1}>
      <Badge color='teal' leftSection={<Avatar size={24} color='green' radius='xl' mr={5} alt='Posted by' />}>User{post.authorId}</Badge>
      <Text>{post.title}</Text>
      <Text color='dimmed' fz='xs'>panel/{post.panelId}</Text>
      <Text fz='sm'>{post.content}</Text>
    </Stack>
  </Paper>
)

export default FeedPost