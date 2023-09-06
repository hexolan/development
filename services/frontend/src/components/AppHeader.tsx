import { Header, Button, Group } from '@mantine/core';
import { Link } from 'react-router-dom';

import panelsLogo from '../assets/logo.svg';

function AppHeader() {
  return (
    <Header height={60} px={20} py={15}>
      <Link to='/'>
        <img src={panelsLogo} height={30} alt='Panels Logo' />
      </Link>

      <Group>
        <Button>Sign In</Button>
      </Group>
    </Header>
  )
}

export default AppHeader;