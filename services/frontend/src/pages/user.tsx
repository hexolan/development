import { Container, Center, Paper, Flex, Avatar, Text, Tabs } from '@mantine/core'
import { IconMessageCircle, IconSettings } from '@tabler/icons-react';
import { useParams } from 'react-router-dom'

type UserPageParams = {
  username: string;
}

function UserPage() {
  const { username } = useParams<UserPageParams>();

  return (
    <Container>
      <Tabs color="teal" radius="md" defaultValue="posts">
        <Paper withBorder shadow='md' radius='md' sx={(theme) => ({ marginTop: theme.spacing.md })}>
          <Flex>
            <Avatar radius="md" size={200} color="lime" />
            <Paper w='100%'>
              <div style={{ position: 'relative', height: '100%' }}>
                <Center h='100%'>
                  <Text weight={600} mr={3}>User:</Text>
                  <Text>{username}</Text>
                </Center>

                <Tabs.List style={{ position: 'absolute', bottom: 0 }}>
                  <Tabs.Tab value="posts" icon={<IconMessageCircle size="0.8rem" />}>Posts</Tabs.Tab>
                  <Tabs.Tab value="about" icon={<IconMessageCircle size="0.8rem" />}>About</Tabs.Tab>
                  <Tabs.Tab value="settings" icon={<IconSettings size="0.8rem" />}>Settings</Tabs.Tab>
                </Tabs.List>
              </div>
            </Paper>
          </Flex>
        </Paper>

        <Tabs.Panel value="posts" pt="xs">
          TODO
        </Tabs.Panel>

        <Tabs.Panel value="about" pt="xs">
          Welcome to my profile.
          TODO
        </Tabs.Panel>

        <Tabs.Panel value="settings" pt="xs">
          TODO: user and admin view only
        </Tabs.Panel>
      </Tabs>
    </Container>
  )
}

export default UserPage