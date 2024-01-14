import { ReactNode } from "react";
import { VStack, Heading, Grid, GridItem } from "@chakra-ui/react";
import { BaseHeader, Footer } from "../../components";

import { useNavigate } from "react-router-dom";

interface AdminBasePageProps {
  children: ReactNode;
  activePage: string;
}

export const AdminBasePage: React.FC<AdminBasePageProps> = ({
  children,
  activePage,
}) => {
  const navigate = useNavigate();
  const pages = ["Dashboard", "Request", "Data", "Users"];

  return (
    <Grid
      templateAreas={`"nav header"
                        "nav main"
                        "nav footer"`}
      gridTemplateRows={"0.2fr 1fr 0.1fr"}
      gridTemplateColumns={"0.2fr 1fr"}
      minH="100svh"
      w="full"
      color="blackAlpha.700"
      fontWeight="bold"
    >
      <GridItem pl="2" area={"header"}>
        <BaseHeader>
          <Heading>{activePage}</Heading>
        </BaseHeader>
      </GridItem>
      <GridItem pl="2" bg="brand_orange.400" area={"nav"}>
        <VStack align="start" pt={8}>
          {pages.map((page) => (
            <Heading
              fontWeight={page === activePage ? "bold" : "normal"}
              onClick={() => navigate(`/admin/${page.toLowerCase()}`)}
              cursor="pointer"
            >
              {page}
            </Heading>
          ))}
        </VStack>
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
