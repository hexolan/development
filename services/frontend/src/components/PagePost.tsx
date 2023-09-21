import { Link } from 'react-router-dom'
import { Paper, Stack, Badge, ThemeIcon, Text } from '@mantine/core'
import { IconUser } from '@tabler/icons-react'

import { useGetUserByIdQuery } from "../app/api/users";
import { Post } from "../app/types/common";

const PagePost = ({ post }: { post: Post }) => {
  const { data } = useGetUserByIdQuery({ id: post.authorId })

  return (
    <Paper shadow="xl" radius="lg" p="lg" withBorder >
      {data && (
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
      )}
      <Stack align="flex-start" mt={2} spacing={1}>
        <Text weight={600}>{post.title}</Text>
        <Text size='sm'>{post.content}</Text>
        <Text size='xs' color='dimmed' mt={3}>Created at | Updated at</Text>
      </Stack>
    </Paper>
  )
}

export default PagePost