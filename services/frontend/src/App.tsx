import { MantineProvider } from '@mantine/core'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'

import AppLayout from './components/AppLayout'
import Homepage from './pages/Home'
import ErrorPage from './pages/Error'

import SignInPage from './pages/SignIn'
import SignUpPage from './pages/SignUp'
import UserPage from './pages/User'
import PanelPage from './pages/Panel'
import ExplorePanelsPage from './pages/ExplorePanels'
import NewPanelPage from './pages/NewPanel'
import PostPage from './pages/Post'

const router = createBrowserRouter([
  {
    element: <AppLayout />,
    errorElement: <AppLayout><ErrorPage /></AppLayout>,
    children: [
      {
        index: true,
        element: <Homepage />,
      },
      {
        path: '/signin',
        element: <SignInPage />,
      },
      {
        path: '/signup',
        element: <SignUpPage />,
      },
      {
        path: '/user/:username',
        element: <UserPage />,
      },
      {
        path: '/panel/:panelName',
        element: <PanelPage />,
      },
      {
        path: '/panels',
        element: <ExplorePanelsPage />,
      },
      {
        path: '/panels/new',
        element: <NewPanelPage />,
      },
      {
        path: '/panel/:panelName/:postId',
        element: <PostPage />,
      }
    ]
  }
])

function App() {
  // todo: change fallback element to Loading component
  // todo: create a loading component
  return (
    <MantineProvider withGlobalStyles withNormalizeCSS>
      <RouterProvider router={router} fallbackElement={<p>Loading...</p>} />
    </MantineProvider>
  );
}

export default App