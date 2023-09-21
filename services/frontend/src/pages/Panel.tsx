import { useOutletContext } from 'react-router-dom'
import { Stack } from '@mantine/core'

import PanelPostFeed from '../components/PanelPostFeed'
import type { PanelContext } from '../components/PanelLayout'

function PanelPage() {
  const { panel } = useOutletContext<PanelContext>()

  return (
    <Stack spacing='md'>
      <PanelPostFeed panelId={panel.id} />
    </Stack>
  )
}

export default PanelPage