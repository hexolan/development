import { Suspense } from 'react';
import { Outlet, useLocation, useNavigate, useParams } from 'react-router-dom'
import { Container, Center, Paper, Flex, Avatar, Text, Tabs } from '@mantine/core'
import { IconMessageCircle, IconAddressBook, IconSettings } from '@tabler/icons-react'

import LoadingBar from './LoadingBar';
import { User } from '../app/types/common';
import { useGetUserByNameQuery } from '../app/api/users';

export type UserContext = {
  user: User
}

type UserPageParams = {
  username: string;
}

function UserLayout() {
  const navigate = useNavigate();
  
  const { pathname } = useLocation()
  let currentTab = 'posts'
  if (pathname.endsWith('about') || pathname.endsWith('about/')) {
    currentTab = 'about'
  } else if (pathname.endsWith('settings') || pathname.endsWith('settings/')) {
    currentTab = 'settings'
  }

  const { username } = useParams<UserPageParams>();
  if (username === undefined) {
    throw Error('username not provided')
  }

  const { data, isLoading } = useGetUserByNameQuery({ username: username })
  if (isLoading) {
    return <LoadingBar />;
  } else if (!data) {
    throw Error('User not found!')
  }

  return (
    <Container mt='xl'>
      <Tabs color='teal' radius='md' value={currentTab} onTabChange={(tab) => navigate(`/user/${data.username}${tab === 'posts' ? '' : `/${tab}`}`)}>
        <Paper withBorder shadow='md' radius='md' sx={(theme) => ({ marginTop: theme.spacing.md })}>
          <Flex>
            <Avatar radius='md' size={200} color='lime' />
            <Paper w='100%'>
              <div style={{ position: 'relative', height: '100%' }}>
                <Center h='100%'>
                  <Text weight={600} mr={3}>User:</Text>
                  <Text>{data.username}</Text>
                </Center>

                <Tabs.List style={{ position: 'absolute', bottom: 0, width: '100%' }} grow>
                  <Tabs.Tab value='posts' icon={<IconMessageCircle size='0.8rem' />}>Posts</Tabs.Tab>
                  <Tabs.Tab value='about' icon={<IconAddressBook size='0.8rem' />}>About</Tabs.Tab>
                  <Tabs.Tab value='settings' icon={<IconSettings size='0.8rem' />}>Settings</Tabs.Tab>
                </Tabs.List>
              </div>
            </Paper>
          </Flex>
        </Paper>
      </Tabs>

      <Suspense>
        <Outlet context={{ user: data } satisfies UserContext} />
      </Suspense>
    </Container>
  )
}

export default UserLayout