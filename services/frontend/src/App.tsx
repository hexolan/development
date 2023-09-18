import { lazy } from 'react'
import { MantineProvider } from '@mantine/core'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'

import AppLayout from './components/AppLayout'
import LoadingBar from './components/LoadingBar'
import ErrorPage from './pages/Error'

const Homepage = lazy(() => import('./pages/Home'))
const SignInPage = lazy(() => import('./pages/SignIn'))
const SignUpPage = lazy(() => import('./pages/SignUp'))
const UserPage = lazy(() => import('./pages/User'))
const ExplorePanelsPage = lazy(() => import('./pages/ExplorePanels'))
const NewPanelPage = lazy(() => import('./pages/NewPanel'))
const PanelPage = lazy(() => import('./pages/Panel'))
const PostPage = lazy(() => import('./pages/Post'))

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
        path: '/panels',
        element: <ExplorePanelsPage />,
      },
      {
        path: '/panels/new',
        element: <NewPanelPage />,
      },
      {
        path: '/panel/:panelName',
        element: <PanelPage />,
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
      <RouterProvider router={router} fallbackElement={<LoadingBar />} />
    </MantineProvider>
  );
}

export default App