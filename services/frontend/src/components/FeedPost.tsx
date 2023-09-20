import { Paper, Stack, Avatar, Badge, Text } from '@mantine/core'

import { useGetUserByIdQuery } from '../app/api/users'
import { useGetPanelByIdQuery } from '../app/api/panels'

import type { Post } from '../app/types/common'

const FeedPost = ({ post }: { post: Post }) => {
  let panelName = '';
  const { panelData, isLoadingPanel } = useGetPanelByIdQuery({ id: post.panelId })
  if (!panelData) {
    panelName = ''
  } else if (isLoadingPanel) {
    panelName = 'Loading...' // todo: spinner
  } else {
    panelName = 'panel/' + panelData.entities[0].name
  }
  
  let authorName = '';
  const { authorData, isLoadingAuthor } = useGetUserByIdQuery({ id: post.authorId })
  if (!authorData) {
    authorName = ''
  } else if (isLoadingAuthor) {
    authorName = 'Loading...'
  } else {
    authorName = authorData.entities[0].username
  }

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