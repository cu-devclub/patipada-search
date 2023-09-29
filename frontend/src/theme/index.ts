import {
  extendTheme,
  theme as base,
  withDefaultColorScheme,
    withDefaultVariant
} from "@chakra-ui/react";
import { mode } from '@chakra-ui/theme-tools';
const inputSelectedStyles = {
    variants: {
        filled: {
            field: {
                _focus: {
                    borderColor:'gray.100',
                    bg: 'gray.100'
                },
                _hover: {
                    borderColor:'gray.100',
                    bg: 'gray.100'
                }
            }
        }
    },
    sizes : {
        md: {
            field: {
                borderRadius: 'none',
            }
        }
    }

}

const brandRing = {
    _focus: {
        ring:2,
        ringColor:'gray.500'
    }
}
const theme = extendTheme(
  {
    sizes : {
        '6xs': '7rem',
        '7xs' : '5rem',
        '8xs' : '4rem'
    },
    fonts: {
      heading: `Montserrat, ${base.fonts?.heading}`,
      body: `Inter ${base.fonts?.body}`,
    },
    components: {
        Button: {
            props: {
                myProp: String, // Define the type for myProp
            },
            variants: {
            primary: (props: { myProp: string }) => ({ // Specify the type for props
                rounded: 'none',
                ...brandRing,
                color: mode('white', 'gray.800')(props),
                backgroundColor: mode('gray.500', 'gray.200')(props),

                _hover: {
                backgroundColor: mode('gray.600', 'gray.300')(props)
                },

                _active: {
                backgroundColor: mode('gray.700', 'gray.400')(props)
                }
            })
            }
        },
    

       Input : {...inputSelectedStyles},
       Select: {...inputSelectedStyles},
       CheckBox: {
        baseStyle: {
            control: {
                borderRadius: 'none',
                ...brandRing
            }
        }
       }
    }
  },
  withDefaultColorScheme({
    colorScheme: "gray",
    components: ["Checkbox"],
  }),
  withDefaultVariant({
    variant: 'filled',
    components: ['Input','Select']
  }),
);

export default theme;
