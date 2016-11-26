
# TODO
* Support for JPEG output
* Custom output file name
* Option to save intermediate images
* Handling of "No Image" images (14:40 often has these)
* Use the current time to nearest 10m
* If using current time and get a 404, roll back 10mins and try again (max rollbacks?)
* Comments
* Remove globals
* Reorder types, functions, vars
* Unit tests
** readImage to test drawing with local files
* Debug logging mode?
* Consistent use of terminology: A tile is drawn to produce an image
* Cropping functionality
** Just crop image before writing
** Improve by skipping downloads for images that aren't needed.
* Create helper for URL construction (unit tested)
* Measure performance
* Percentage completion in-line?
* Download speeds
* Summarise output image: location, file size, dimensions, format, cropping?
* --help override? If possible, add examples.
* Reorder flags
