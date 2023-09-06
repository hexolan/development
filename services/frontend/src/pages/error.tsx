import { Title, Text, Button, Center, Container, Group, rem } from '@mantine/core';
import { Link, useRouteError, isRouteErrorResponse } from 'react-router-dom';

const ErrorPage = () => {
  const error = useRouteError();

  let title: string = 'Uh, oh!'
  let subTitle: string = 'Something went wrong.'
  if (isRouteErrorResponse(error)) {
    title = `Error ${error.status}`
    subTitle = error.statusText
  } else if (error instanceof Error) {
    subTitle = error.message
  }

  return (
    <Center h='100%'>
      <Container>
        <Title
          sx={(theme) => ({
            textAlign: 'center',
            fontWeight: 900,
            fontSize: rem(38),
            [theme.fn.smallerThan('sm')]: {
              fontSize: rem(32),
            },
          })}
        >
          {title}
          </Title>
        <Text 
          size='lg'
          color='dimmed'
          sx={(theme) => ({
            maxWidth: rem(500),
            textAlign: 'center',
            marginTop: theme.spacing.xl,
            marginBottom: `calc(${theme.spacing.xl} * 1.5)`,
          })}
        >
          {subTitle}
        </Text>

        <Group position='center'>
          <Button component={Link} to='/' variant='subtle' size='md'>
            Back to Home
          </Button>
        </Group>
      </Container>
    </Center>
  )
}

export default ErrorPage;