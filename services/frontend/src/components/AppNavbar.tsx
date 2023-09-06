import { createElement } from 'react';

import { Navbar, ThemeIcon, Group, Text, UnstyledButton,rem } from '@mantine/core';
import { IconTrendingUp, IconSearch, IconMessages } from '@tabler/icons-react';
import { NavLink } from "react-router-dom";

const NavbarButton = ({ text, page, icon }: { text: string, page: string, icon: any }) => (
  <NavLink to={page} style={{ textDecoration: "none" }}>
    {({ isActive }) => (
      <UnstyledButton
        sx={(theme) => ({
          display: "block",
          width: "100%",
          padding: theme.spacing.xs,
          borderRadius: theme.radius.sm,

          backgroundColor: isActive ? theme.colors.gray[0] : "inherit",
          "&:hover": {
            backgroundColor: theme.colors.gray[0],
          },
        })}
      >
        <Group>
          <ThemeIcon color="teal" variant="light">
            { createElement(icon, {size: "1rem"}) }
          </ThemeIcon>

          <Text size="sm">{text}</Text>
        </Group>
      </UnstyledButton>
    )}
  </NavLink>
)

function AppNavbar() {
  return (
    <Navbar width={{ base: 300 }} p="xs">
      <Navbar.Section py="xs">
        <Text
          size="xs"
          color="dimmed"
          sx={(theme) => ({
            marginLeft: theme.spacing.xs,
            marginBottom: theme.spacing.xs,
            fontWeight: 500,
          })}
        >
          Browse
        </Text>
        <NavbarButton text="Feed" page="/" icon={IconTrendingUp} />
        <NavbarButton text="Find Panels" page="/panels" icon={IconSearch} />
      </Navbar.Section>

      <Navbar.Section 
        grow 
        sx={(theme) => ({
          borderTop: `${rem(1)} solid ${theme.colors.gray[3]}`,
          paddingTop: theme.spacing.xs
        })}
      >
        <Text 
          size="xs"
          color="dimmed"
          sx={(theme) => ({
            margin: theme.spacing.xs,
            fontWeight: 500,
          })}
        >
          Subscribed Panels
        </Text>
        <NavbarButton text="panel/Panel" page="/panel/Panel" icon={IconMessages} />
      </Navbar.Section>
    </Navbar>
  )
}

export default AppNavbar;