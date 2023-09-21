import { Paper, Stack, Avatar, Badge, Text } from '@mantine/core'

import { useGetUserByIdQuery } from '../app/api/users'
import { useGetPanelByIdQuery } from '../app/api/panels'

import type { Post } from '../app/types/common'

const FeedPost = ({ post }: { post: Post }) => {
  const getPanelElement = () => {
    const { data, isLoading } = useGetPanelByIdQuery({ id: post.panelId })
    if (isLoading) {
      return <Text color='dimmed' size='xs'>Loading...</Text> // todo: change this
    } else if (!data) {
      return <Text color='red' size='xs'>Error Loading Panel Data</Text> // todo: change this
    } else {
      return <Text color='dimmed' size='xs'>{'panel/' + data.name}</Text>
    }
  }

  const getAuthorElement = () => {
    const { data, isLoading } = useGetUserByIdQuery({ id: post.authorId })
    if (isLoading) {
      return null
    } else if (!data) {
      <Text color='red' size='xs'>Error Loading Author Data</Text> // todo: change this
    } else {
      return (
        <Badge 
          pl={0} 
          color='teal'
          leftSection={<Avatar size={24} color='green' radius='xl' mr={5} alt={'Created by ' + data.username} />}
        >
          {'user/' + data.username}
        </Badge>
      )
    }
  }

  const panelElement = getPanelElement()
  const authorElement = getAuthorElement()

  return (
    <Paper shadow="xl" radius="lg" p="lg" withBorder>
      <Stack align="flex-start" spacing={1}>
        {authorElement}
        <Text>{post.title}</Text>
        {panelElement}
        <Text size='sm' truncate>{post.content}</Text>
      </Stack>
    </Paper>
  )
}

export default FeedPost