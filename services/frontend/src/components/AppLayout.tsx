import { AppShell } from "@mantine/core";
import { Outlet } from "react-router-dom";

import AppNavbar from "./AppNavbar";
import AppHeader from "./AppHeader";

function AppLayout(props: any) {
  return (
    <AppShell
      navbar={<AppNavbar />} 
      header={<AppHeader />}
      padding="md"
    >
      {props?.children ? props.children : <Outlet /> }
    </AppShell>
  );
}

export default AppLayout;