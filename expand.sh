for i in apple banana cherry
do
  cat my-jobs-tmpl.yaml | sed "s/\$ITEM/$i/" > ./jobs/job-$i.yaml
done
