
-- This file is formated by datatool, DO NOT EDIT THIS FILE.
-- $id$
module("Module")
md5sum = md5sum or {}
md5sum.Field="19bbb5343a1ae7014d1f6fac734fa4b0"
Field=
{ -- 
    [2]={ -- 2
        base={ -- 2.base
            number=1,
            float=0.1,
            boolean=true,
            string={ -- 2.base.string
                empty="",
                char="123",
                lf="123\n123",
                crlf="123\n\r",
                space=" 123",
                dquote="\"123\"",
                tab="\t123",
            },
        },
        mixin={ -- 2.mixin
            [2]="number key",
            key="string key",
            [0.2]="float key",
        },
        array={ -- 2.array
            number={ -- 2.array.number
                multi_line={ -- 2.array.number.multi_line
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
                single_line={ -- total: 13
                    1,2,3,4,5,1,2,3,4,5,1,2,3,
                },
            },
            float={ -- 2.array.float
                multi_line={ -- total: 26
                    0.1,0.2,0.3,0.4,0.5,0.1,0.2,0.3,0.4,0.5,0.1,
                    0.2,0.3,0.1,0.2,0.3,0.4,0.5,0.1,0.2,0.3,0.4,
                    0.5,0.1,0.2,0.3,
                },
                single_line={ -- total: 13
                    0.1,0.2,0.3,0.4,0.5,0.1,0.2,0.3,0.4,0.5,0.1,
                    0.2,0.3,
                },
            },
            unorder={ -- 2.array.unorder
                [-1]=true,
                [10]=false,
                [21]="text1",
                [82]=1,
                [88]=0.2,
            },
            boolean={ -- 2.array.boolean
                multi_line={ -- total: 15
                     true, true, true, true,false, true, true,
                    false, true, true,false, true, true, true,
                     true,
                },
                single_line={ -- total: 8
                    false,false, true, true, true, true, true,
                    false,
                },
            },
            string={ -- 2.array.string
                multi_line={ -- total: 19
                     "\t123\t",    "text",    "text",     "kkk",
                     "123\r\n",    "text",   " text", "\"123\"",
                        "text","123\n123",    "text",    "text",
                        "text", "123\n\r",    "text",    "text",
                        "text","123\n123",        "",
                },
                single_line={ -- total: 8
                      "\t123",  "123\n","123\n\r",   "text",  " text",
                      "\"123",   "text",    "123",
                },
            },
        },
        table={ -- 2.table
            number=1,
            float=0.1,
            boolean=true,
            string={ -- 2.table.string
                empty="",
                char="123",
                lf="123\n123",
                crlf="123\n\r",
                space=" 123",
                dquote="\"123\"",
                tab="\t123",
            },
        },
    },
}
		