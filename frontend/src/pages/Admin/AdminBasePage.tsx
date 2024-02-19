import { ReactNode } from "react";
import { Grid, GridItem } from "@chakra-ui/react";
import {
  AdminSidebar,
  EditRequestHeader,
  Footer,
} from "../../components/layout";

interface AdminBasePageProps {
  children: ReactNode;
  activePage: string;
  requestID?: string;
}

export const AdminBasePage: React.FC<AdminBasePageProps> = ({
  children,
  activePage,
  requestID,
}) => {
  return (
    <Grid
      templateAreas={`"nav header"
                        "nav main"
                        "nav footer"`}
      gridTemplateRows={"0.2fr 2fr 0.2fr"}
      gridTemplateColumns={"0.2fr 1fr"}
      h="100svh"
      w="full"
      color="blackAlpha.700"
      fontWeight="bold"
    >
      <GridItem pl="2" area={"header"}>
        <EditRequestHeader activePage={activePage} requestID={requestID} />
      </GridItem>
      <GridItem area={"nav"}>
        <AdminSidebar activePage={activePage} />
      </GridItem>
      <GridItem pl="2" area={"main"}>
        {children}
      </GridItem>
      <GridItem area={"footer"}>
        <Footer />
      </GridItem>
    </Grid>
  );
};
