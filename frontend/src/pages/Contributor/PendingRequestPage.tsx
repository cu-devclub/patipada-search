import { Grid } from "@chakra-ui/react";
import { PendingRequestHeader } from "../../components/contributor/pending-request";
import { Footer } from "../../components";
function PendingRequestPage() {
  return (
    <Grid templateRows="auto 1fr auto" gap={4} w="full" minH="100svh">
      <PendingRequestHeader />
      {/* //TODO : pending request table */}
      <Footer />
    </Grid>
  );
}

export default PendingRequestPage;
