import { Link } from 'react-router-dom'
import { Paper, Box, Stack, Badge, Text, Group, ThemeIcon } from '@mantine/core'
import { IconUser, IconMessages } from '@tabler/icons-react'

import { useGetUserByIdQuery } from '../app/api/users'
import { useGetPanelByIdQuery } from '../app/api/panels'
import type { Post } from '../app/types/common'

const FeedPost = ({ post, hidePanel, hideAuthor }: { post: Post, hidePanel?: boolean, hideAuthor?: boolean }) => {
  // fetch panel info
  let panelElement = <></>
  const { data: panelData } = useGetPanelByIdQuery({ id: post.panelId })
  if (!hidePanel) {
    if (!panelData) {
      panelElement = <Text color='red' size='xs'>Error Loading Panel Data</Text>
    } else {
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
          to={`/panel/${panelData.name}`}
        >
          {`panel/${panelData.name}`}
        </Badge>
      )
    }
  }
  
  // fetch author info
  let authorElement = <></>
  const { data: authorData } = useGetUserByIdQuery({ id: post.authorId })
  if (!hideAuthor) {
    if (!authorData) {
      authorElement = <Text color='red' size='xs'>Error Loading Author Data</Text>
    } else {
      authorElement = (
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
      <Group spacing='xs'>
        {panelElement}
        {authorElement}
      </Group>
      <Stack align='flex-start' mt={2} spacing={1}>
        <Box component={Link} to={panelData ? `/panel/${panelData.name}/post/${post.id}` : '#'} style={{ textDecoration: 'none', color: 'inherit' }}>
          <Text weight={600} lineClamp={1}>{post.title}</Text>
          <Text size='sm' lineClamp={2}>{post.content}</Text>
          <Text size='xs' color='dimmed' mt={3}>Click to View</Text>
        </Box>
      </Stack>
    </Paper>
  )
}

export default FeedPost