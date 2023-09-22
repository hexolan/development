import { Link } from 'react-router-dom'
import { Paper, Box, Stack, Badge, Text, ThemeIcon } from '@mantine/core'
import { IconUser, IconMessages } from '@tabler/icons-react'

import { useGetUserByIdQuery } from '../app/api/users'
import { useGetPanelByIdQuery } from '../app/api/panels'

import type { Post, Panel } from '../app/types/common'

const FeedPost = ({ post, panelInfo }: { post: Post, panelInfo: Panel | undefined }) => {
  const loadPanelData = () => {
    if (panelInfo) {
      return {
        panelElement: null, // element is hidden (as the post is on a panel page)
        panelName: panelInfo.name
      }
    }

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
      return {
        panelElement: (
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
        ),
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

  // todo: wireframe loaders
  // todo: show createdAt timestamp? (not priority)
  const { panelElement, panelName } = loadPanelData()
  const { authorElement } = loadAuthorData()
  
  return (
    <Paper shadow='xl' radius='lg' p='lg' withBorder>
      {panelElement} {authorElement} 
      <Stack align='flex-start' mt={2} spacing={1}>
        <Box component={Link} to={panelName ? `/panel/${panelName}/post/${post.id}` : '#'} style={{ textDecoration: 'none', color: 'inherit' }}>
          <Text weight={600} lineClamp={1}>{post.title}</Text>
          <Text size='sm' lineClamp={2}>{post.content}</Text>
          <Text size='xs' color='dimmed' mt={3}>Click to View</Text>
        </Box>
      </Stack>
    </Paper>
  )
}

export default FeedPost