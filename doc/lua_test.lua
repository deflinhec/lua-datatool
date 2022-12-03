-- $id$
module("Module")
md5sum = md5sum or {}
md5sum.Field="a55862bc45004f9af169868018799c2c"
Field=
{ -- Module.Field
    [2]={ -- Module.Field[2]
        base={ -- Module.Field[2].base
            string={ -- Module.Field[2].base.string
                empty="",
                dquote="\"123\"",
                space=" 123",
                lf="123\n123",
                crlf="123\n\r",
                tab="\t123",
                char="123",
            },
            number=1,
            boolean=true,
            float=0.1,
        },
        table={ -- Module.Field[2].table
            string={ -- Module.Field[2].table.string
                empty="",
                dquote="\"123\"",
                space=" 123",
                lf="123\n123",
                crlf="123\n\r",
                tab="\t123",
                char="123",
            },
            number=1,
            boolean=true,
            float=0.1,
        },
        array={ -- Module.Field[2].array
            string={ -- Module.Field[2].array.string
                single_line={ -- total: 8
                    "\t123","123\n","123\n\r","text"," text","\"123","text","123",
                },
                multi_line={ -- total: 19
                     "\t123\t",    "text",    "text",     "kkk",
                     "123\r\n",    "text",   " text", "\"123\"",
                        "text","123\n123",    "text",    "text",
                        "text", "123\n\r",    "text",    "text",
                        "text","123\n123",        "",
                },
            },
            number={ -- Module.Field[2].array.number
                single_line={ -- total: 13
                    1,2,3,4,5,1,2,3,4,5,1,2,3,
                },
                multi_line={ -- Module.Field[2].array.number.multi_line
                    case1={ -- total: 30
                              1,      2,      3,      4,      5,
                              6,      7,      8,      9,     10,
                             11,     12,      2,     77,      4,
                              5,    100,     21,      5,      1,
                             23,      3,      4,      5,      1,
                              2,      3,      1,      2,1000000,
                    },
                    case2={ -- total: 30
                          1,  2,  3,  4,  5,  6,  7,  8,  9, 10,
                         11, 12,  2, 77,  4,  5,100, 21,  5, 44,
                         23,  3,  4,  5,  1,  2,  3,  1,  2,100,
                    },
                },
            },
            boolean={ -- Module.Field[2].array.boolean
                single_line={ -- total: 8
                    false,false,true,true,true,true,true,false,
                },
                multi_line={ -- total: 15
                     true, true, true, true,false, true, true,
                    false, true, true,false, true, true, true,
                     true,
                },
            },
            float={ -- Module.Field[2].array.float
                single_line={ -- total: 13
                    0.1,0.2,0.3,0.4,0.5,0.1,0.2,0.3,0.4,0.5,0.1,0.2,0.3,
                },
                multi_line={ -- total: 26
                    0.1,0.2,0.3,0.4,0.5,0.1,0.2,0.3,0.4,0.5,0.1,
                    0.2,0.3,0.1,0.2,0.3,0.4,0.5,0.1,0.2,0.3,0.4,
                    0.5,0.1,0.2,0.3,
                },
            },
            unorder={ -- Module.Field[2].array.unorder
                [-1]=true,
                [10]=false,
                [21]="text1",
                [82]=1,
                [88]=0.2,
            },
        },
        mixin={ -- Module.Field[2].mixin
            [0.2]="float key",
            [2]="number key",
            key="string key",
        },
    },
}