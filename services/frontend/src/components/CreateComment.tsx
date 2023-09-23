import { Paper, Group, Box, ThemeIcon, Text } from '@mantine/core'
import { IconWriting } from '@tabler/icons-react'

import type { Post } from "../app/types/common"

const CreateComment = ({ post }: { post: Post }) => (
  <Paper shadow='sm' radius='md' p='md' mb='sm' withBorder>
    <Group>
      <ThemeIcon color='teal' variant='light' size='xl'><IconWriting /></ThemeIcon>
      <Box>
        <Text size='sm'>Create Comment ({post.id})</Text>
        <Text color='dimmed' size='xs' mt={3}>Todo</Text>
      </Box>
    </Group>
  </Paper>
)

export default CreateComment