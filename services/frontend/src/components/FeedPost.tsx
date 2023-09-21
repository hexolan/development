import { Link } from 'react-router-dom'
import { Paper, Stack, Avatar, Badge, Text } from '@mantine/core'

import { useGetUserByIdQuery } from '../app/api/users'
import { useGetPanelByIdQuery } from '../app/api/panels'

import type { Post } from '../app/types/common'

const FeedPost = ({ post }: { post: Post }) => {
  const loadPanelData = () => {
    const { data, isLoading } = useGetPanelByIdQuery({ id: post.panelId })
    if (isLoading) {
      return {
        panelElement: <Text color='dimmed' size='xs'>Loading...</Text>, // todo: change this
        panelName: undefined
      }
    } else if (!data) {
      return {
        panelElement: <Text color='red' size='xs'>Error Loading Panel Data</Text>, // todo: change this
        panelName: null
      }
    } else {
      return {
        panelElement: <Text color='dimmed' size='xs'>{'panel/' + data.name}</Text>,
        panelName: data.name
      }
    }
  }

  const loadAuthorData = () => {
    const { data, isLoading } = useGetUserByIdQuery({ id: post.authorId })
    if (isLoading) {
      return {
        authorElement: null,
        authorName: undefined
      }
    } else if (!data) {
      return {
        authorElement: <Text color='red' size='xs'>Error Loading Author Data</Text>, // todo: change this
        authorName: null
      }
    } else {
      return {
        authorElement: (
          <Badge 
            pl={0}
            color='teal'
            leftSection={<Avatar size={24} color='green' radius='xl' mr={5} alt={`Created by ${data.username}`} />}
            component={Link}
            to={`/user/${data.username}`}
          >
            {`user/${data.username}`}
          </Badge>
        ),
        authorName: data.username
      }
    }
  }

  const { panelElement, panelName } = loadPanelData()
  const { authorElement } = loadAuthorData()

  return (
    <Paper shadow="xl" radius="lg" p="lg" withBorder component={Link} to={panelName ? `/panel/${panelName}/post/${post.id}` : '#'}>
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