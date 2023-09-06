import { Text } from '@mantine/core';
import { useParams } from "react-router-dom";

type PanelPageParams = {
  panelName: string;
}

function PanelPage() {
  // TODO: using data loaders for panel
  // https://stackoverflow.com/questions/75740241/what-type-to-declare-for-useloaderdata-and-where
  const { panelName } = useParams<PanelPageParams>();

  return (
    <Text>Panel - {panelName}</Text>
  )
}

export default PanelPage;