import { Suspense } from 'react'
import { Outlet, useParams } from 'react-router-dom'
import { Paper, Container, Text, rem } from '@mantine/core'

import LoadingBar from '../components/LoadingBar'
import { useGetPanelByNameQuery } from '../app/api/panels'
import type { Panel } from '../app/types/common'

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

  const { data, isLoading } = useGetPanelByNameQuery({ name: panelName })
  if (isLoading) {
    return <LoadingBar />;
  } else if (!data) {
    // todo: extract err msg
    throw Error('Panel not found!')
  }

  return (
    <>
      <Paper px="xl" py={rem(50)} shadow='md' sx={{ borderBottom: '1px' }}>
        <Text size='lg'>{data.name}</Text>
        <Text size='sm' color='dimmed'>{data.description}</Text>

        {/* todo: new panel button here (on right side, opposite side, of two elems above */}
      </Paper>
      <Container mt='xl'>
        <Suspense fallback={<LoadingBar />}>
          <Outlet context={{panel: data}} />
        </Suspense>
      </Container>
    </>
  )
}

export default PanelLayout