# Changelog

All notable changes to this project will be documented in this file.
This project adheres to Semantic Versioning.

## [1.6.1] 2021-05-18

### Added

- Dump command: generate core dumps from within Delve (@aarzilli)
- Toggle command: toggle breakpoints on or off (@alexsaezm)
- DAP server improvements (@polinasok, @hyangah, @suzmue)
- Delve now parses and uses the .eh_frame section when available (@aarzilli)
- Add linespec argument to 'continue' command (@icholy)
- Add optional format argument to 'print' and 'display' commands (@aarzilli)

### Fixed

- Fixed file reference handling with DWARF5 compilation units (@thanm)
- Fix embedded field searching (@aarzilli)
- Fix off by one error reading core files (@aarzilli)
- Correctly read G address on linux/arm64
- Avoid double removal of temp built binary (@polinasok)
- Fix temp binary deletion race in DAP server (@polinasok)
- Fix shutdown related bugs in RPC server (@aarzilli)
- Fix crashes induced by RequestManualStop (@aarzilli)
- Fix handling of DW_OP_piece (@aarzilli)
- Correctly truncate the result of binary operations on integers (@aarzilli)

### Changed

- Dropped Go 1.13 support (@aarzilli)
- Improved documentation (@ChrisHines, @aarzilli, @becheran, @hedenface, @andreimatei, @ehershey , @hyangah)
- Allow examinememory to use an expression (@felixge)
- Improve gdb server check on newer ARM based macs (@oxisto)
- CPU registers can now be used in expressions (@aarzilli)
- DAP: Add type information to variables (@suzmue)
- DAP: Support setting breakpoints while target is running (@polinasok)
- DAP: Implement function breakpoints (@suzmue)

## [1.6.0] 2021-01-28

### Added

- Support for debugging darwin/arm64 (i.e. macOS on Apple Silicon) (#2285, @oxisto)
- Support for debugging Go1.16 (#2214, @aarzilli)
- DAP: support for attaching to a local process (#2260, @polinasok)
- DAP: fill the `evaluateName` field when returning variables, enabling "Add to Watch" and "Copy as Expression" features of VSCode (#2292, @polinasok)
- Added WaitSince, WaitReason to `service/api.Goroutine` and to the `goroutines` command (#2264, #2283, #2270, @dlsniper, @nd, @aarzilli)
- Syntax highlighting for Go code (#2294, @aarzilli)
- Added flag `CallReturn` to `service/api.Thread` to distinguish return values filled by a `stepOut` command from the ones filled by a `call` command (#2230, @aarzilli)

### Fixed

- Fix occasional "Access is denied" error when debugging on Windows (#2281, @nd)
- Register formatting on ARM64 (#2289, @dujinze)
- Miscellaneous bug fixes (#2232, #2255, #2280, #2286, #2291, #2309, #2293, @aarzilli, @polinasok, @hitzhangjie)

### Changed

- The `goroutines` command can be interrupted by pressing Ctrl-C (#2278, @aarzilli)
- Using a TeamCity instance provided by JetBrains for Continuous Integration (#2298, #2307, #2311, #2315, #2326, @artspb, @nd, @aarzilli, @derekparker)
- Improvements to documentation and error messages (#2266, #2265, #2273, #2299, @andreimatei, @hitzhangjie, @zamai, @polinasok)
