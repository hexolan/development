import { ReactNode, Suspense } from 'react'
import { AppShell, Progress } from '@mantine/core'
import { Outlet } from 'react-router-dom'

import AppNavbar from './AppNavbar'
import AppHeader from './AppHeader'

interface AppLayoutProps {
  children?: ReactNode;
}

function AppLayout(props: AppLayoutProps) {
  return (
    <AppShell
      navbar={<AppNavbar />} 
      header={<AppHeader />}
      padding={0}
    >
      <Suspense fallback={<Progress color="lime" radius="xs" value={100} striped animate />}>
        {props?.children ? props.children : <Outlet /> }
      </Suspense>
    </AppShell>
  );
}

export default AppLayout