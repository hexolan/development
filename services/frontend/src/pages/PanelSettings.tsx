import { Paper, Text, Button } from '@mantine/core'

import { useAppSelector } from '../app/hooks'

function PanelSettingsPage() {
  // Check permissions
  const currentUser = useAppSelector((state) => state.auth.currentUser)
  if (!currentUser || !currentUser.isAdmin) {
    throw new Error('You do not have permission to view that page.')
  }

  return (
    <Paper mt='md' radius='lg' shadow='md' p='lg' withBorder>
      <Button color='teal'>Modify Panel</Button>
      <Button color='red'>Delete Panel</Button>
      <Text color='red'>Error: error message</Text>
    </Paper>
  )
}

export default PanelSettingsPage