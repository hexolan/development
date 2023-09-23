import { Link } from 'react-router-dom'
import { Paper, Box, Stack, Badge, Text, ThemeIcon } from '@mantine/core'
import { IconUser, IconMessages } from '@tabler/icons-react'

import { useGetUserByIdQuery } from '../app/api/users'
import { useGetPanelByIdQuery } from '../app/api/panels'
import type { Post, Panel, User } from '../app/types/common'

const FeedPost = ({ post, panelInfo, authorInfo }: { post: Post, panelInfo?: Panel, authorInfo?: User }) => {
  // todo: if I have to call useGetPanelByIdQuery and useGetUserByIdQuery each time,
  // change the above panelInfo and authorInfo to hidePanel and hideAuthor bools and clean up mess below
  
  // Load Panel Attributes
  const panelAttrs = { element: <></>, name: ''}
  const { data: panelData, isLoading: panelIsLoading } = useGetPanelByIdQuery({ id: post.panelId })
  if (!panelInfo) {
    if (panelIsLoading) {
      panelAttrs.element = <></>
    } else if (!panelData) {
      panelAttrs.element = <Text color='red' size='xs'>Error Loading Panel Data</Text>
    } else {
      panelAttrs.name = panelData.name
      panelAttrs.element = (
        <Badge
          pl={0}
          color='orange'
          leftSection={
            <ThemeIcon color='orange' size={24} radius='xl' mr={5}>
              <IconMessages size={12} />
            </ThemeIcon>
          }
          component={Link}
          to={`/panel/${panelData.name}`}
        >
          {`panel/${panelData.name}`}
        </Badge>
      )
    }
  } else {
    panelAttrs.name = panelInfo.name
  }
  
  // Load Author Attributes
  const authorAttrs = { element: <></> }
  const { data: authorData, isLoading: authorIsLoading } = useGetUserByIdQuery({ id: post.authorId })
  if (!authorInfo) {
    if (authorIsLoading) {
      authorAttrs.element = <></>
    } else if (!authorData) {
      authorAttrs.element = <Text color='red' size='xs'>Error Loading Author Data</Text>
    } else {
      authorAttrs.element = (
        <Badge
          pl={0}
          color='teal'
          leftSection={
            <ThemeIcon color='teal' size={24} radius='xl' mr={5}>
              <IconUser size={12} />
            </ThemeIcon>
          }
          component={Link}
          to={`/user/${authorData.username}`}
        >
          {`user/${authorData.username}`}
        </Badge>
      )
    }
  }


  // todo: wireframe loaders
  // todo: show createdAt timestamp? (not priority)
  return (
    <Paper shadow='xl' radius='lg' p='lg' withBorder>
      {panelAttrs.element} {authorAttrs.element} 
      <Stack align='flex-start' mt={2} spacing={1}>
        <Box component={Link} to={panelAttrs.name ? `/panel/${panelAttrs.name}/post/${post.id}` : '#'} style={{ textDecoration: 'none', color: 'inherit' }}>
          <Text weight={600} lineClamp={1}>{post.title}</Text>
          <Text size='sm' lineClamp={2}>{post.content}</Text>
          <Text size='xs' color='dimmed' mt={3}>Click to View</Text>
        </Box>
      </Stack>
    </Paper>
  )
}

export default FeedPost