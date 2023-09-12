import { MantineProvider } from '@mantine/core'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'

import AppLayout from './components/AppLayout'
import IndexPage from './pages/index'
import SignInPage from './pages/signIn'
import SignUpPage from './pages/signUp'
import UserPage from './pages/user'
import PanelPage from './pages/panel'
import ExplorePanelsPage from './pages/explorePanels'
import NewPanelPage from './pages/newPanel'
import PostPage from './pages/post'
import ErrorPage from './pages/error'

const router = createBrowserRouter([
  {
    element: <AppLayout />,
    errorElement: <AppLayout><ErrorPage /></AppLayout>,
    children: [
      {
        index: true,
        element: <IndexPage />,
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
  return (
    <MantineProvider withGlobalStyles withNormalizeCSS>
      <RouterProvider router={router} fallbackElement={<p>Loading...</p>} />
    </MantineProvider>
  );
}

export default App