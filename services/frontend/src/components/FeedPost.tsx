import { Link } from 'react-router-dom'
import { Paper, Stack, Badge, Text, ThemeIcon } from '@mantine/core'
import { IconUser, IconMessages } from '@tabler/icons-react'

import { useGetUserByIdQuery } from '../app/api/users'
import { useGetPanelByIdQuery } from '../app/api/panels'

import type { Post } from '../app/types/common'

const FeedPost = ({ post, hidePanel }: { post: Post, hidePanel: boolean | undefined }) => {
  // todo: wireframe loaders
  // todo: show createdAt timestamp? (not priority)
  const loadPanelData = () => {
    const { data, isLoading } = useGetPanelByIdQuery({ id: post.panelId })
    if (isLoading) {
      return {
        panelElement: null,
        panelName: undefined
      }
    } else if (!data) {
      return {
        panelElement: <Text color='red' size='xs'>Error Loading Panel Data</Text>, // todo: change this
        panelName: null
      }
    } else {
      let panelElement = null
      if (!hidePanel) {
        panelElement = (
          <Badge
            pl={0}
            color='orange'
            leftSection={
              <ThemeIcon color='orange' size={24} radius='xl' mr={5}>
                <IconMessages size={12} />
              </ThemeIcon>
            }
            component={Link}
            to={`/panel/${data.name}`}
          >
            {`panel/${data.name}`}
          </Badge>
        )
      }

      return {
        panelElement: panelElement,
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
      // leftSection={<Avatar size={24} color='green' radius='xl' mr={5} alt={`Created by ${data.username}`} />}
      return {
        authorElement: (
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
        ),
        authorName: data.username
      }
    }
  }

  const { panelElement, panelName } = loadPanelData()
  const { authorElement } = loadAuthorData()

  return (
    <Paper shadow="xl" radius="lg" p="lg" withBorder component={Link} to={panelName ? `/panel/${panelName}/post/${post.id}` : '#'}>
      {panelElement} {authorElement} 
      <Stack align="flex-start" mt={2} spacing={1}>
        <Text lineClamp={1}>{post.title}</Text>
        <Text size='sm' color='dimmed' lineClamp={2}>{post.content}</Text>
      </Stack>
    </Paper>
  )
}

export default FeedPost