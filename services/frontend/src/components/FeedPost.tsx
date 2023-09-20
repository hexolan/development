import { Paper, Stack, Avatar, Badge, Text } from '@mantine/core'

import { useGetUserByIdQuery } from '../app/api/users'
import { useGetPanelByIdQuery } from '../app/api/panels'

import type { Post } from '../app/types/common'

const FeedPost = ({ post }: { post: Post }) => {
  const getPanelName = () => {
    const { data, isLoading } = useGetPanelByIdQuery({ id: post.panelId })
    if (isLoading) {
      return 'Loading...'
    } else if (!data) {
      return 'Error loading panel'
    } else {
      return data.name
    }
  }

  const getAuthorName = () => {
    const { data, isLoading } = useGetUserByIdQuery({ id: post.authorId })
    if (isLoading) {
      return 'Loading...'
    } else if (!data) {
      return 'Error loading author' // todo
    } else {
      return data.username
    }
  }

  const panelName = getPanelName()
  const authorName = getAuthorName()

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