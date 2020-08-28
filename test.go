//use any name you want, but should be unique
package androidbp1_test                                                                                                                                

import (
        "android/soong/android"
        "android/soong/cc"
        "fmt" //For log print
)


func init() {
	fmt.Println("first call when compile");
    android.RegisterModuleType("androidbp_test1_defaults", testDefaultsFactory)
}

func testDefaultsFactory() android.Module {
        module := cc.DefaultsFactory()
        android.AddLoadHook(module, testDefaults)
        return module
}

func testDefaults(ctx android.LoadHookContext) {
        type props struct {
                Srcs []string
                Cflags []string
                Shared_libs []string
                Header_libs []string  //local module compile include dirs
                Export_header_lib_headers []string //if this is a lib, other module compile depend on this module, should export the headers
        }
        p := &props{}

        p.Cflags = append(p.Cflags, "-Werror")
        p.Cflags = append(p.Cflags, "-Wall")

        vars := ctx.Config().VendorConfig("mtkPlugin")

        fmt.Println("MYTEST1_SUPPORT support:", vars.Bool("MYTEST1_SUPPORT"))
        fmt.Println("MYTEST1_DEBUG_SUPPORT support:", vars.Bool("MYTEST1_DEBUG_SUPPORT"))
        fmt.Println("MYTEST2_DEBUG_SUPPORT support:", vars.Bool("MYTEST2_DEBUG_SUPPORT"))


        ret_mytest := vars.String("MYTEST1_SUPPORT")
        if ret_mytest == "yes" {
        	fmt.Println("mytest1 support")
        }

        if vars.Bool("MYTEST1_SUPPORT") {
                p.Srcs = append(p.Srcs, ":mytest1_Sources")

                p.Shared_libs = append(p.Shared_libs, "libutils")

                p.Cflags = append(p.Cflags, "-Wno-unused-parameter")
                p.Cflags = append(p.Cflags, "-Wno-unused-function")
                if vars.Bool("MYTEST1_DEBUG_SUPPORT") {
        			//add this for debug config
        			p.Cflags = append(p.Cflags, "-DMYTESTDEBUG_1")
        		}

                p.Header_libs = append(p.Header_libs, "mytest1_headers")
                p.Export_header_lib_headers = append(p.Export_header_lib_headers, "mytest1_headers")
        } else {
	        p.Srcs = append(p.Srcs, ":mytest2_Sources")

	        p.Shared_libs = append(p.Shared_libs, "libutils")

	        p.Cflags = append(p.Cflags, "-Wno-unused-parameter")
	        p.Cflags = append(p.Cflags, "-Wno-unused-function")
	        if vars.Bool("MYTEST2_DEBUG_SUPPORT") {
				//add this for debug config
				p.Cflags = append(p.Cflags, "-DMYTEST2DEBUG_2")
			}

	        p.Header_libs = append(p.Header_libs, "mytest2_headers")
	        p.Export_header_lib_headers = append(p.Export_header_lib_headers, "mytest2_headers")
        }

        ctx.AppendProperties(p)
}
