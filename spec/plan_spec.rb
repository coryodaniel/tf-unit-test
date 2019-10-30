require 'json'

RSpec.describe "planning" do
  let(:plan) { File.read("./plan.json") }
  let(:tf) { json = JSON.parse(plan) }

  def get_changed_attributes(path)
    resource = get_changed_resource(path)
    resource ? resource["changedAttributes"] : nil
  end

  def get_changed_resource(path)
    return nil if tf["changedResources"].empty?

    tf["changedResources"].find {|resource| resource["path"] == path}
  end

  it "successfully creates a plan" do
    expect(tf["errors"]).to be_empty
  end

  it "registers a temp file path" do
    attrs = get_changed_attributes("module.mymod.local_file.foo")
    expect(attrs["filename"]["new"]["value"]).to eq("/tmp/index.md")
  end
end
