import { Paper, Flex, Textarea, ActionIcon } from '@mantine/core'
import { IconWriting } from '@tabler/icons-react'

import type { Post } from "../app/types/common"

const CreateComment = ({ post }: { post: Post }) => (
  <Paper shadow='sm' radius='md' p='md' mb='sm' withBorder>
    <Flex gap='sm' align='center' direction='row' wrap='nowrap'>
      <Textarea size='xs' radius='lg' w='100%' variant='filled' placeholder='Input comment...' />
      
      <ActionIcon radius='lg' color='teal' variant='outline' size='xl' aria-label='Post Comment'>
        <IconWriting />
      </ActionIcon>
    </Flex>
  </Paper>
)

export default CreateComment