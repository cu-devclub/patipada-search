import React, { ReactNode } from "react";
import { UseRadioProps, useRadio, Box, Circle } from "@chakra-ui/react";

interface RadioCardProps extends UseRadioProps {
    children: ReactNode;
}

const RadioCard: React.FC<RadioCardProps> = (props) => {
    const { getInputProps, getCheckboxProps } = useRadio(props);

    const input = getInputProps();
    const checkbox = getCheckboxProps();

    return (
        <Box as="label">
            <input {...input} />
            <Circle
                {...checkbox}
                cursor="pointer"
                bg="#262f38"
                boxShadow="md"
                _checked={{
                    bg: "#fc7613",
                    color: "white",
                    borderColor: "#fc7613",
                }}
                _hover={{
                    bg: "#7c8799",
                }}
                px={{ base: 4, md: 6, lg: 6 }}
                py={{ base: 2, md: 4, lg: 4 }}
                mx={{ base: 0.5, md: 2, lg: 2 }}
                my={{ base: 2, md: 4, lg: 4 }}
            >
                {props.children}
            </Circle>
        </Box>
    );
};

export default RadioCard;