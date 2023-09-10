import { Link } from 'react-router-dom'
import { Container, Stack, Paper, Title, Text, Anchor, Divider, ThemeIcon, Group } from '@mantine/core'
import { IconMessages, IconTableOff } from '@tabler/icons-react'

const ExplorePanelsPage = () => {
  return (
    <Container mt={15}>
      <Title>Explore Panels</Title>
      <Text color='dimmed' size='sm' mt={5}>
        Alternatively you could <Anchor size='sm' component={Link} to='/panels/new'>create your own.</Anchor>
      </Text>
      <Divider my="md" variant="dotted" />

      <Stack spacing='sm' align='stretch'>
        <Paper shadow="xl" radius="md" p="md" withBorder component={Link} to='/panel/Panel'>
          <Group>
            <ThemeIcon color='teal' variant='light' size='xl'><IconMessages /></ThemeIcon>
            <div>
              <Text weight={600}>Panel</Text>
              <Text>The first and ergo de facto primary panel.</Text>
              <Text color='dimmed' size='xs' mt={3}>Click to View</Text>
            </div>
          </Group>
        </Paper>

        <Paper shadow="xl" radius="md" p="md" withBorder component={Link} to='/'>
          <Group>
            <ThemeIcon color='red' variant='light' size='xl'><IconTableOff /></ThemeIcon>
            <div>
              <Text weight={600}>Note</Text>
              <Text>This page is exemplary as this feature is currently unimplemented.</Text>
              <Text color='dimmed' size='xs' mt={3}>Planned Functionality</Text>
            </div>
          </Group>
        </Paper>
      </Stack>
    </Container>
  )
}

export default ExplorePanelsPage