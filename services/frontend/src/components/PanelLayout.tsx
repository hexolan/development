import { Suspense } from 'react'
import { Link, Outlet, useParams } from 'react-router-dom'
import { Paper, Container, Group, Box, Text, Button, rem } from '@mantine/core'

import LoadingBar from '../components/LoadingBar'
import { useAppSelector } from '../app/hooks'
import { useGetPanelByNameQuery } from '../app/api/panels'
import type { Panel } from '../app/types/common'
import type { ErrorResponse } from '../app/types/api'

export type PanelContext = {
  panel: Panel
}

type PanelParams = {
  panelName: string
}

function PanelLayout() {
  const { panelName } = useParams<PanelParams>();
  if (panelName === undefined) {
    throw Error('panel name not provided')
  }
  
  let newPostButton = null
  const currentUser = useAppSelector((state) => state.auth.currentUser)
  const { data, error, isLoading } = useGetPanelByNameQuery({ name: panelName })
  if (isLoading) {
    return <LoadingBar />;
  } else if (!data) {
    if (!error) {
      throw Error('Unknown error occured')
    } else if ('data' in error) {
      const errResponse = error.data as ErrorResponse
      if (errResponse.msg) {
        throw Error(errResponse.msg)
      } else {
        throw Error('Unexpected API error occured')
      }
    } else {
      throw Error('Failed to access the API')
    }
  } else if (currentUser) {
    // panel info was succesfully fetched.
    // provide a 'new post' button if the user is authenticated
    newPostButton = <Button size='xs' variant="filled" color="teal" component={Link} to={`/panel/${data.name}/posts/new`}>Create Post</Button>
  }
  
  return (
    <>
      <Paper py={rem(50)} shadow='md' sx={{ borderBottom: '1px' }}>
        <Container>
          <Group position='apart'>
            <Box>
              <Text size='lg'>{data.name}</Text>
              <Text size='sm' color='dimmed'>{data.description}</Text>
            </Box>

            {newPostButton}
          </Group>
        </Container>
      </Paper>
      <Container mt='xl'>
        <Suspense>
          <Outlet context={{panel: data} satisfies PanelContext} />
        </Suspense>
      </Container>
    </>
  )
}

export default PanelLayout