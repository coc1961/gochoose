# gochoose

## Example

```sh
Select yes or no in shell script

>option=`echo -e "yes\nno" | gochoose`; echo Option Selected $option

>Option Selected yes



select and change dir interactively

cd `ls -d -1 */ | gochoose`



```