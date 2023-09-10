import { Center, Container, Paper, Title, TextInput, Textarea, Button } from '@mantine/core'

const NewPanelPage = () => {
  return (
    <Center h='95%'>
      <Container w='25%'>
        <Title align='center' weight={900}>Create a Panel</Title>

        <Paper withBorder shadow='md' radius='md' p={30} mt={30}>
          <TextInput label='Name' placeholder="e.g. music, programming, football" />
          <Textarea label='Description' placeholder="e.g. The place to talk about all things music related..." mt={6} />
          <Button variant='outline' color='teal' fullWidth mt='xl'>Create it!</Button>
        </Paper>
      </Container>
    </Center>
  )
}

export default NewPanelPage