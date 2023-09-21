import { useParams } from 'react-router-dom'
import { Paper, Container, Stack, Text, rem } from '@mantine/core'

import LoadingBar from '../components/LoadingBar'
import PanelPostFeed from '../components/PanelPostFeed'
import { useGetPanelByNameQuery } from '../app/api/panels'

type PanelPageParams = {
  panelName: string;
}

function PanelPage() {
  const { panelName } = useParams<PanelPageParams>();
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
        <Stack spacing='md'>
          <PanelPostFeed panelId={data.id} />
        </Stack>
      </Container>
    </>
  )
}

export default PanelPage